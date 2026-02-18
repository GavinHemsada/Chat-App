package websocket

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
)

type BroadcastMessage struct {
	RoomID     string
	Message    []byte
	UserID     string // For processing incoming messages
	FromRedis  bool   // Flag to prevent republishing messages from Redis
}

type MessageProcessor func(*BroadcastMessage) *BroadcastMessage

type Hub struct {
	clients         map[string]map[*Client]bool // roomID -> clients
	broadcast       chan *BroadcastMessage
	register        chan *Client
	unregister      chan *Client
	mu              sync.RWMutex
	redisPub        *redis.Client
	redisSub        *redis.PubSub
	messageProcessor MessageProcessor
}

func NewHub(redisClient *redis.Client, processor MessageProcessor) *Hub {
	return &Hub{
		clients:          make(map[string]map[*Client]bool),
		broadcast:        make(chan *BroadcastMessage, 256),
		register:         make(chan *Client),
		unregister:       make(chan *Client),
		redisPub:         redisClient,
		messageProcessor: processor,
	}
}

func (h *Hub) Run() {
	ctx := context.Background()
	
	// Subscribe to Redis for messages from other server instances (if Redis is available)
	if h.redisPub != nil {
		// Subscribe to a pattern that matches all room channels
		h.redisSub = h.redisPub.PSubscribe(ctx, "chat:room:*")
		// Start Redis listener in goroutine
		go h.listenRedis(ctx)
		log.Println("Redis pub/sub initialized for WebSocket")
	}
    
    for {
        select {
        case client := <-h.register:
            h.mu.Lock()
            if h.clients[client.roomID] == nil {
                h.clients[client.roomID] = make(map[*Client]bool)
            }
            h.clients[client.roomID][client] = true
            h.mu.Unlock()
            log.Printf("Client registered to room %s", client.roomID)
            
        case client := <-h.unregister:
            h.mu.Lock()
            if clients, ok := h.clients[client.roomID]; ok {
                if _, ok := clients[client]; ok {
                    delete(clients, client)
                    close(client.send)
                    if len(clients) == 0 {
                        delete(h.clients, client.roomID)
                    }
                }
            }
            h.mu.Unlock()
            log.Printf("Client unregistered from room %s", client.roomID)
            
        case message := <-h.broadcast:
            // Process message if processor is available
            if h.messageProcessor != nil && message.UserID != "" {
                processedMsg := h.messageProcessor(message)
                if processedMsg != nil {
                    message = processedMsg
                }
            }
            
            // Send to local clients first
            h.sendToLocalClients(message)
            
            // Publish to Redis for other server instances (if Redis is available)
            // Only publish if:
            // 1. Message is processed (UserID is empty means it's ready to broadcast)
            // 2. Message didn't come from Redis (FromRedis is false)
            if h.redisPub != nil && message.UserID == "" && !message.FromRedis {
                // Serialize the full BroadcastMessage for Redis
                messageBytes, err := json.Marshal(message)
                if err != nil {
                    log.Printf("Error marshaling message for Redis: %v", err)
                } else {
                    // Use room-specific channel to avoid cross-room message leakage
                   	channel := "chat:room:" + message.RoomID
                   	if err := h.redisPub.Publish(ctx, channel, messageBytes).Err(); err != nil {
                   		log.Printf("Error publishing to Redis: %v", err)
                   	} else {
                   		log.Printf("Published message to Redis channel: %s", channel)
                   	}
                }
            }
        }
    }
}

func (h *Hub) sendToLocalClients(message *BroadcastMessage) {
    h.mu.RLock()
    defer h.mu.RUnlock()
    
    if clients, ok := h.clients[message.RoomID]; ok {
        for client := range clients {
            select {
            case client.send <- message.Message:
            default:
                close(client.send)
                delete(clients, client)
            }
        }
    }
}

func (h *Hub) listenRedis(ctx context.Context) {
	if h.redisSub == nil {
		return
	}
	
	ch := h.redisSub.Channel()
	
	for msg := range ch {
		// Extract room ID from channel name (format: chat:room:{roomID})
		// For PSubscribe, msg.Channel contains the actual channel name that matched the pattern
		prefix := "chat:room:"
		if !strings.HasPrefix(msg.Channel, prefix) {
			log.Printf("Unexpected Redis channel format: %s", msg.Channel)
			continue
		}
		roomID := msg.Channel[len(prefix):]
		if roomID == "" {
			log.Printf("Empty room ID in channel: %s", msg.Channel)
			continue
		}
		
		var broadcastMsg BroadcastMessage
		if err := json.Unmarshal([]byte(msg.Payload), &broadcastMsg); err != nil {
			log.Printf("Error unmarshaling Redis message: %v", err)
			continue
		}
		
		// Ensure RoomID matches (safety check)
		if broadcastMsg.RoomID != roomID {
			broadcastMsg.RoomID = roomID
		}
		
		// Mark as from Redis to prevent republishing
		broadcastMsg.FromRedis = true
		broadcastMsg.UserID = "" // Ensure it's not processed again
		
		// Don't process again, just broadcast to local clients
		// This message came from another server instance and was already processed there
		log.Printf("Received message from Redis for room: %s", roomID)
		h.sendToLocalClients(&broadcastMsg)
	}
}