package controllers

// GetStatus controller
func GetStatus() (int, map[string]string, error) {
	body := make(map[string]string)
	body["status"] = "ok"

	log.Info("Handling request for /status")
	status, body := 200, body
	return status, body, nil
}
