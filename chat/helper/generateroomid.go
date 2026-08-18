package helper

func GenerateRoomID(ip1, ip2 string) string {
	if ip1 < ip2 {
		return ip1 + "_" + ip2
	}
	return ip2 + "_" + ip1
}
