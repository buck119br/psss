package net

const (
	TCPI_OPT_TIMESTAMPS = 1
	TCPI_OPT_SACK       = 2
	TCPI_OPT_WSCALE     = 4
	TCPI_OPT_ECN        = 8  /* ECN was negociated at TCP session init */
	TCPI_OPT_ECN_SEEN   = 16 /* we received at least one packet with ECT */
	TCPI_OPT_SYN_DATA   = 32 /* SYN-ACK acked data in SYN sent or rcvd */
)
