package main

const (
	TCPI_OPT_TIMESTAMPS = 1
	TCPI_OPT_SACK       = 2
	TCPI_OPT_WSCALE     = 4
	TCPI_OPT_ECN        = 8  /* ECN was negociated at TCP session init */
	TCPI_OPT_ECN_SEEN   = 16 /* we received at least one packet with ECT */
	TCPI_OPT_SYN_DATA   = 32 /* SYN-ACK acked data in SYN sent or rcvd */
)

type TCPInfo struct {
	State           uint8
	Ca_state        uint8
	Retransmits     uint8
	Probes          uint8
	Backoff         uint8
	Options         uint8
	Pad_cgo_0       [2]byte
	Rto             uint32
	Ato             uint32
	Snd_mss         uint32
	Rcv_mss         uint32
	Unacked         uint32
	Sacked          uint32
	Lost            uint32
	Retrans         uint32
	Fackets         uint32
	Last_data_sent  uint32
	Last_ack_sent   uint32
	Last_data_recv  uint32
	Last_ack_recv   uint32
	Pmtu            uint32
	Rcv_ssthresh    uint32
	Rtt             uint32
	Rttvar          uint32
	Snd_ssthresh    uint32
	Snd_cwnd        uint32
	Advmss          uint32
	Reordering      uint32
	Rcv_rtt         uint32
	Rcv_space       uint32
	Total_retrans   uint32
	Pacing_rate     uint64
	Max_pacing_rate uint64
	Bytes_acked     uint64 /* RFC4898 tcpEStatsAppHCThruOctetsAcked */
	Bytes_received  uint64 /* RFC4898 tcpEStatsAppHCThruOctetsReceived */
	Segs_out        uint32 /* RFC4898 tcpEStatsPerfSegsOut */
	Segs_in         uint32 /* RFC4898 tcpEStatsPerfSegsIn */
	Notsent_bytes   uint32
	Min_rtt         uint32
	Data_segs_in    uint32 /* RFC4898 tcpEStatsDataSegsIn */
	Data_segs_out   uint32 /* RFC4898 tcpEStatsDataSegsOut */
	Delivery_rate   uint64
	Busy_time       uint64 /* Time (usec) busy sending data */
	Rwnd_limited    uint64 /* Time (usec) limited by receive window */
	Sndbuf_limited  uint64 /* Time (usec) limited by send buffer */
}
