package main

/*import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Function to fetch network I/O stats from /proc/net/dev
func getNetworkIO() (map[string]map[string]uint64, error) {
	// Open the /proc/net/dev file
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Map to hold network statistics
	networkStats := make(map[string]map[string]uint64)

	// Read each line in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip headers and empty lines
		if strings.Contains(line, "Inter-") || strings.TrimSpace(line) == "" {
			continue
		}

		// Extract the interface name and the network statistics
		parts := strings.Fields(line)
		if len(parts) < 17 {
			continue
		}

		interfaceName := strings.TrimSpace(parts[0][:len(parts[0])-1]) // Remove trailing colon
		receivedBytes := parts[1]
		transmittedBytes := parts[9]

		// Convert string values to uint64
		received, err := strToUint64(receivedBytes)
		if err != nil {
			return nil, err
		}

		transmitted, err := strToUint64(transmittedBytes)
		if err != nil {
			return nil, err
		}

		// Store the data in the map
		networkStats[interfaceName] = map[string]uint64{
			"received_bytes":    received,
			"transmitted_bytes": transmitted,
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return networkStats, nil
}

// Helper function to convert string to uint64
func strToUint64(str string) (uint64, error) {
	var value uint64
	_, err := fmt.Sscanf(str, "%d", &value)
	return value, err
}

func main() {
	// Continuous polling every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Poll indefinitely
	for {
		select {
		case <-ticker.C: // Every 5 seconds, fetch network stats
			// Get the network I/O stats
			networkStats, err := getNetworkIO()
			if err != nil {
				log.Printf("Error fetching network stats: %v", err)
				continue
			}

			// Print network stats for each interface
			fmt.Println("Network I/O Stats:")
			for iface, stats := range networkStats {
				fmt.Printf("Interface: %s\n", iface)
				fmt.Printf("  Received Bytes: %d\n", stats["received_bytes"])
				fmt.Printf("  Transmitted Bytes: %d\n", stats["transmitted_bytes"])
			}
		}
	}
}
*/
