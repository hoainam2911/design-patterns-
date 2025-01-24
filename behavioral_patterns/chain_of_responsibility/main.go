package main

import (
	"fmt"
)

// Handler interface declares a method for building the chain of handlers and a method for executing a request.
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request map[string]interface{}) string
}

type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.nextHandler = handler
	return handler
}

func (h *BaseHandler) Handle(request map[string]interface{}) string {
	if h.nextHandler != nil {
		return h.nextHandler.Handle(request)
	}
	return ""
}

type AuthenticationHandler struct {
	BaseHandler
}

func (h *AuthenticationHandler) Handle(request map[string]interface{}) string {
	if !request["authenticated"].(bool) {
		return "Authentication failed."
	}
	fmt.Println("Authentication succeeded.")
	return h.BaseHandler.Handle(request)
}

type AdminPermissionHandler struct {
	BaseHandler
}

func (h *AdminPermissionHandler) Handle(request map[string]interface{}) string {
	if request["is_admin"].(bool) {
		fmt.Println("Admin permissions granted.")
	} else {
		fmt.Println("Standard user permissions granted.")
	}
	return h.BaseHandler.Handle(request)
}

type DataValidationHandler struct {
	BaseHandler
}

func (h *DataValidationHandler) Handle(request map[string]interface{}) string {
	if request["data"] == nil {
		return "Invalid data."
	}
	fmt.Println("Data is valid.")
	return h.BaseHandler.Handle(request)
}

type RateLimitingHandler struct {
	BaseHandler
}

func (h *RateLimitingHandler) Handle(request map[string]interface{}) string {
	if request["rate_limited"].(bool) {
		return "Rate limit exceeded."
	}
	fmt.Println("Rate limit check passed.")
	return h.BaseHandler.Handle(request)
}

type CacheHandler struct {
	BaseHandler
	cache map[string]string
}

func (h *CacheHandler) Handle(request map[string]interface{}) string {
	key := request["data"].(string)
	if value, exists := h.cache[key]; exists {
		return fmt.Sprintf("Cache hit: %s", value)
	}
	fmt.Println("Cache miss.")
	return h.BaseHandler.Handle(request)
}

func main() {
	// Simulated cache
	cacheStore := map[string]string{
		"order123": "Cached order details",
	}

	// Creating handlers
	authHandler := &AuthenticationHandler{}
	permissionHandler := &AdminPermissionHandler{}
	validationHandler := &DataValidationHandler{}
	rateLimitingHandler := &RateLimitingHandler{}
	cacheHandler := &CacheHandler{cache: cacheStore}

	// Building the chain
	authHandler.SetNext(permissionHandler).
		SetNext(validationHandler).
		SetNext(rateLimitingHandler).
		SetNext(cacheHandler)

	// Example request
	request := map[string]interface{}{
		"authenticated": true,
		"is_admin": true,
		"data": "order123",
		"rate_limited": false,
	}

	// Pass the request through the chain
	result := authHandler.Handle(request)
	if result != "" {
		fmt.Println(result)
	}
}
