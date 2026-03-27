package scanner

import "strings"

// DetectService identifies the probable service running on a port
func DetectService(port int, banner string) string {
	banner = strings.ToLower(banner)
	switch {
	case strings.Contains(banner, "ssh"):
		return "SSH"
	case strings.Contains(banner, "ftp"):
		return "FTP"
	case strings.Contains(banner, "smtp"):
		return "SMTP"
	case strings.Contains(banner, "imap"):
		return "IMAP"
	case strings.Contains(banner, "pop3"):
		return "POP3"
	case strings.Contains(banner, "http"):
		return "HTTP"
	case strings.Contains(banner, "https"):
		return "HTTPS"
	case strings.Contains(banner, "mysql"):
		return "MySQL"
	case strings.Contains(banner, "postgres"):
		return "PostgreSQL"
	default:
		switch port {
		case 22:
			return "SSH"
		case 21:
			return "FTP"
		case 25:
			return "SMTP"
		case 80:
			return "HTTP"
		case 443:
			return "HTTPS"
		case 3306:
			return "MySQL"
		default:
			return "Unknown"
		}
	}
}