// XiXiOrchard/internal/monitoring/monitoring.go
package monitoring

import (
	"fmt"
	"log"
	"net/smtp"
)

type Monitor struct {
	EmailAlertsEnabled bool
	EmailRecipient     string
	SmtpServer         string
	SmtpPort           int
	SmtpUser           string
	SmtpPassword       string
}

func NewMonitor(enableEmail bool, recipient, smtpServer, smtpUser, smtpPassword string, smtpPort int) *Monitor {
	return &Monitor{
		EmailAlertsEnabled: enableEmail,
		EmailRecipient:     recipient,
		SmtpServer:         smtpServer,
		SmtpPort:           smtpPort,
		SmtpUser:           smtpUser,
		SmtpPassword:       smtpPassword,
	}
}

// SendAlert sends an email alert with the given subject and message.
func (m *Monitor) SendAlert(subject, message string) {
	if m.EmailAlertsEnabled {
		err := smtp.SendMail(
			fmt.Sprintf("%s:%d", m.SmtpServer, m.SmtpPort),
			smtp.PlainAuth("", m.SmtpUser, m.SmtpPassword, m.SmtpServer),
			m.SmtpUser,
			[]string{m.EmailRecipient},
			[]byte("Subject: "+subject+"\r\n\r\n"+message),
		)
		if err != nil {
			log.Printf("Failed to send email alert: %v", err)
		} else {
			log.Println("Email alert sent successfully")
		}
	} else {
		log.Println("Email alerts are disabled")
	}
}

// CalculateMaxDrawdown calculates the maximum drawdown of the strategy results.
func CalculateMaxDrawdown(prices []float64) float64 {
	peak := prices[0]
	maxDrawdown := 0.0

	for _, price := range prices {
		if price > peak {
			peak = price
		}
		drawdown := (peak - price) / peak
		if drawdown > maxDrawdown {
			maxDrawdown = drawdown
		}
	}

	return maxDrawdown
}

// CalculateProfitRatio calculates the profit ratio of the strategy results.
func CalculateProfitRatio(prices []float64) float64 {
	initial := prices[0]
	final := prices[len(prices)-1]
	return (final - initial) / initial
}

// MonitorResults monitors the strategy results and triggers alerts based on conditions.
func (m *Monitor) MonitorResults(prices []float64) {
	maxDrawdown := CalculateMaxDrawdown(prices)
	profitRatio := CalculateProfitRatio(prices)

	log.Printf("Max Drawdown: %.2f%%, Profit Ratio: %.2f%%", maxDrawdown*100, profitRatio*100)

	if maxDrawdown > 0.1 {
		m.SendAlert("High Drawdown Alert", "The strategy experienced a drawdown of more than 10%.")
	}

	if profitRatio < 0 {
		m.SendAlert("Negative Profit Alert", "The strategy ended with a negative profit.")
	}
}
