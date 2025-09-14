package broker

import (
	"bytes"
	"fmt"
	"log"

	"mqt-tech-broker/services"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)

type CustomAuthHook struct {
	*auth.Hook // Embed the base Auth hook for default behavior if needed
}

type CustomLoggerOptions struct {
	Server *mqtt.Server
}

type CustomLogger struct {
	mqtt.HookBase
	Server *mqtt.Server
}

func (h *CustomLogger) ID() string {
	return "customLogger"
}

func (h *CustomLogger) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnSubscribed,
		mqtt.OnUnsubscribed,
		mqtt.OnPublished,
		mqtt.OnPublish,
	}, []byte{b})
}

func (h *CustomAuthHook) OnConnectAuthenticate(cl *mqtt.Client, pk packets.Packet) bool {
	username := string(pk.Connect.Username)
	password := string(pk.Connect.Password)

	if services.ValidateDevice(username, password) {
		log.Printf("✅ Client %s authenticated successfully.", cl.ID)
		return true
	}

	log.Printf("❌ Client %s authentication failed.", cl.ID)
	return false
}

func Start() {
	// Create new MQTT server
	server := mqtt.New(nil)

	// Initialize custom auth hook with base auth hook
	authHook := &CustomAuthHook{
		Hook: &auth.Hook{},
	}

	err := server.AddHook(authHook, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 MQTT Broker Starting")
	// Listener TCP di port 1883
	tcp := listeners.NewTCP(listeners.Config{ID: "t1", Address: ":1883"})
	errtcp := server.AddListener(tcp)
	if errtcp != nil {
		log.Fatal(errtcp)
	}

	log.Println("🚀 Websocket Starting")
	ws := listeners.NewTCP(listeners.Config{ID: "ws1", Address: ":8883"})
	errws := server.AddListener(ws)
	if errws != nil {
		log.Fatal(errws)
	}

	err = server.AddHook(new(CustomLogger), &CustomLoggerOptions{
		Server: server,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 MQTT Broker running on :1883")
	log.Println("🚀 MQTT Websocket running on :8883")
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

	select {}
}

func (h *CustomLogger) OnConnect(cl *mqtt.Client, pk packets.Packet) error {
	h.Log.Info("client connected", "client", cl.ID)
	return nil
}

func (h *CustomLogger) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	if err != nil {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire, "error", err)
	} else {
		h.Log.Info("client disconnected", "client", cl.ID, "expire", expire)
	}

}

func (h *CustomLogger) OnSubscribed(cl *mqtt.Client, pk packets.Packet, reasonCodes []byte) {
	h.Log.Info(fmt.Sprintf("subscribed qos=%v", reasonCodes), "client", cl.ID, "filters", pk.Filters)
}

func (h *CustomLogger) OnUnsubscribed(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info("unsubscribed", "client", cl.ID, "filters", pk.Filters)
}

func (h *CustomLogger) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	h.Log.Info("received from client",
		"client", cl.ID,
		"topic", pk.TopicName,
		"payload", string(pk.Payload),
	)

	// Simpan ke database
	services.LogMessage(
		pk.TopicName,       // Topic sebagai VirtualPinID
		string(pk.Payload), // Payload isi value
	)

	pkx := pk

	return pkx, nil
}

func (h *CustomLogger) OnPublished(cl *mqtt.Client, pk packets.Packet) {
	h.Log.Info("published to client", "client", cl.ID, "payload", string(pk.Payload))
}
