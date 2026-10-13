package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	const envFile = ".env"
	const keyName = "JWT_SECRET"

	secret := make([]byte, 128)
	if _, err := rand.Read(secret); err != nil {
		panic(err)
	}
	newKey := base64.StdEncoding.EncodeToString(secret)

	content := ""
	if data, err := os.ReadFile(envFile); err == nil {
		content = string(data)
	}

	lines := strings.Split(content, "\n")
	found := false

	for i, line := range lines {
		if strings.HasPrefix(line, keyName+"=") {
			lines[i] = fmt.Sprintf("%s=%s", keyName, newKey)
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", keyName, newKey))
	}

	output := strings.Join(lines, "\n")
	if err := os.WriteFile(envFile, []byte(output), 0644); err != nil {
		panic(err)
	}

	fmt.Println("JWT_SECRET updated successfully 🔐")
}
