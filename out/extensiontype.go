package godicttls

// source: https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values-1.csv
// last updated: March 2023

const (
	Extension Name uint16 = 0
	server_name uint16 = 0
	max_fragment_length uint16 = 1
	client_certificate_url uint16 = 2
	trusted_ca_keys uint16 = 3
	truncated_hmac uint16 = 4
	status_request uint16 = 5
	user_mapping uint16 = 6
	client_authz uint16 = 7
	server_authz uint16 = 8
	cert_type uint16 = 9
	supported_groups (renamed from "elliptic_curves") uint16 = 10
	ec_point_formats uint16 = 11
	srp uint16 = 12
	signature_algorithms uint16 = 13
	use_srtp uint16 = 14
	heartbeat uint16 = 15
	application_layer_protocol_negotiation uint16 = 16
	status_request_v2 uint16 = 17
	signed_certificate_timestamp uint16 = 18
	client_certificate_type uint16 = 19
	server_certificate_type uint16 = 20
	padding uint16 = 21
	encrypt_then_mac uint16 = 22
	extended_master_secret uint16 = 23
	token_binding uint16 = 24
	cached_info uint16 = 25
	tls_lts uint16 = 26
	compress_certificate uint16 = 27
	record_size_limit uint16 = 28
	pwd_protect uint16 = 29
	pwd_clear uint16 = 30
	password_salt uint16 = 31
	ticket_pinning uint16 = 32
	tls_cert_with_extern_psk uint16 = 33
	delegated_credentials uint16 = 34
	session_ticket (renamed from "SessionTicket TLS") uint16 = 35
	TLMSP uint16 = 36
	TLMSP_proxying uint16 = 37
	TLMSP_delegate uint16 = 38
	supported_ekt_ciphers uint16 = 39
	Reserved uint16 = 40
	pre_shared_key uint16 = 41
	early_data uint16 = 42
	supported_versions uint16 = 43
	cookie uint16 = 44
	psk_key_exchange_modes uint16 = 45
	Reserved uint16 = 46
	certificate_authorities uint16 = 47
	oid_filters uint16 = 48
	post_handshake_auth uint16 = 49
	signature_algorithms_cert uint16 = 50
	key_share uint16 = 51
	transparency_info uint16 = 52
	connection_id (deprecated) uint16 = 53
	connection_id uint16 = 54
	external_id_hash uint16 = 55
	external_session_id uint16 = 56
	quic_transport_parameters uint16 = 57
	ticket_request uint16 = 58
	dnssec_chain uint16 = 59
	Unassigned uint16 = 0
	Reserved uint16 = 2570
	Unassigned uint16 = 0
	Reserved uint16 = 6682
	Unassigned uint16 = 0
	Reserved uint16 = 10794
	Unassigned uint16 = 0
	Reserved uint16 = 14906
	Unassigned uint16 = 0
	Reserved uint16 = 19018
	Unassigned uint16 = 0
	Reserved uint16 = 23130
	Unassigned uint16 = 0
	Reserved uint16 = 27242
	Unassigned uint16 = 0
	Reserved uint16 = 31354
	Unassigned uint16 = 0
	Reserved uint16 = 35466
	Unassigned uint16 = 0
	Reserved uint16 = 39578
	Unassigned uint16 = 0
	Reserved uint16 = 43690
	Unassigned uint16 = 0
	Reserved uint16 = 47802
	Unassigned uint16 = 0
	Reserved uint16 = 51914
	Unassigned uint16 = 0
	Reserved uint16 = 56026
	Unassigned uint16 = 0
	Reserved uint16 = 60138
	Unassigned uint16 = 0
	Reserved uint16 = 64250
	Unassigned uint16 = 0
	Reserved for Private Use uint16 = 65280
	renegotiation_info uint16 = 65281
	Reserved for Private Use uint16 = 0
)

var Uint16MapToString = map[uint16]string{
	0: "Extension Name",
	0: "server_name",
	1: "max_fragment_length",
	2: "client_certificate_url",
	3: "trusted_ca_keys",
	4: "truncated_hmac",
	5: "status_request",
	6: "user_mapping",
	7: "client_authz",
	8: "server_authz",
	9: "cert_type",
	10: "supported_groups (renamed from \"elliptic_curves\")",
	11: "ec_point_formats",
	12: "srp",
	13: "signature_algorithms",
	14: "use_srtp",
	15: "heartbeat",
	16: "application_layer_protocol_negotiation",
	17: "status_request_v2",
	18: "signed_certificate_timestamp",
	19: "client_certificate_type",
	20: "server_certificate_type",
	21: "padding",
	22: "encrypt_then_mac",
	23: "extended_master_secret",
	24: "token_binding",
	25: "cached_info",
	26: "tls_lts",
	27: "compress_certificate",
	28: "record_size_limit",
	29: "pwd_protect",
	30: "pwd_clear",
	31: "password_salt",
	32: "ticket_pinning",
	33: "tls_cert_with_extern_psk",
	34: "delegated_credentials",
	35: "session_ticket (renamed from \"SessionTicket TLS\")",
	36: "TLMSP",
	37: "TLMSP_proxying",
	38: "TLMSP_delegate",
	39: "supported_ekt_ciphers",
	40: "Reserved",
	41: "pre_shared_key",
	42: "early_data",
	43: "supported_versions",
	44: "cookie",
	45: "psk_key_exchange_modes",
	46: "Reserved",
	47: "certificate_authorities",
	48: "oid_filters",
	49: "post_handshake_auth",
	50: "signature_algorithms_cert",
	51: "key_share",
	52: "transparency_info",
	53: "connection_id (deprecated)",
	54: "connection_id",
	55: "external_id_hash",
	56: "external_session_id",
	57: "quic_transport_parameters",
	58: "ticket_request",
	59: "dnssec_chain",
	0: "Unassigned",
	2570: "Reserved",
	0: "Unassigned",
	6682: "Reserved",
	0: "Unassigned",
	10794: "Reserved",
	0: "Unassigned",
	14906: "Reserved",
	0: "Unassigned",
	19018: "Reserved",
	0: "Unassigned",
	23130: "Reserved",
	0: "Unassigned",
	27242: "Reserved",
	0: "Unassigned",
	31354: "Reserved",
	0: "Unassigned",
	35466: "Reserved",
	0: "Unassigned",
	39578: "Reserved",
	0: "Unassigned",
	43690: "Reserved",
	0: "Unassigned",
	47802: "Reserved",
	0: "Unassigned",
	51914: "Reserved",
	0: "Unassigned",
	56026: "Reserved",
	0: "Unassigned",
	60138: "Reserved",
	0: "Unassigned",
	64250: "Reserved",
	0: "Unassigned",
	65280: "Reserved for Private Use",
	65281: "renegotiation_info",
	0: "Reserved for Private Use",
}

var StringToUint16 = map[string]uint16{
	"Extension Name": 0,
	"server_name": 0,
	"max_fragment_length": 1,
	"client_certificate_url": 2,
	"trusted_ca_keys": 3,
	"truncated_hmac": 4,
	"status_request": 5,
	"user_mapping": 6,
	"client_authz": 7,
	"server_authz": 8,
	"cert_type": 9,
	"supported_groups (renamed from \"elliptic_curves\")": 10,
	"ec_point_formats": 11,
	"srp": 12,
	"signature_algorithms": 13,
	"use_srtp": 14,
	"heartbeat": 15,
	"application_layer_protocol_negotiation": 16,
	"status_request_v2": 17,
	"signed_certificate_timestamp": 18,
	"client_certificate_type": 19,
	"server_certificate_type": 20,
	"padding": 21,
	"encrypt_then_mac": 22,
	"extended_master_secret": 23,
	"token_binding": 24,
	"cached_info": 25,
	"tls_lts": 26,
	"compress_certificate": 27,
	"record_size_limit": 28,
	"pwd_protect": 29,
	"pwd_clear": 30,
	"password_salt": 31,
	"ticket_pinning": 32,
	"tls_cert_with_extern_psk": 33,
	"delegated_credentials": 34,
	"session_ticket (renamed from \"SessionTicket TLS\")": 35,
	"TLMSP": 36,
	"TLMSP_proxying": 37,
	"TLMSP_delegate": 38,
	"supported_ekt_ciphers": 39,
	"Reserved": 40,
	"pre_shared_key": 41,
	"early_data": 42,
	"supported_versions": 43,
	"cookie": 44,
	"psk_key_exchange_modes": 45,
	"Reserved": 46,
	"certificate_authorities": 47,
	"oid_filters": 48,
	"post_handshake_auth": 49,
	"signature_algorithms_cert": 50,
	"key_share": 51,
	"transparency_info": 52,
	"connection_id (deprecated)": 53,
	"connection_id": 54,
	"external_id_hash": 55,
	"external_session_id": 56,
	"quic_transport_parameters": 57,
	"ticket_request": 58,
	"dnssec_chain": 59,
	"Unassigned": 0,
	"Reserved": 2570,
	"Unassigned": 0,
	"Reserved": 6682,
	"Unassigned": 0,
	"Reserved": 10794,
	"Unassigned": 0,
	"Reserved": 14906,
	"Unassigned": 0,
	"Reserved": 19018,
	"Unassigned": 0,
	"Reserved": 23130,
	"Unassigned": 0,
	"Reserved": 27242,
	"Unassigned": 0,
	"Reserved": 31354,
	"Unassigned": 0,
	"Reserved": 35466,
	"Unassigned": 0,
	"Reserved": 39578,
	"Unassigned": 0,
	"Reserved": 43690,
	"Unassigned": 0,
	"Reserved": 47802,
	"Unassigned": 0,
	"Reserved": 51914,
	"Unassigned": 0,
	"Reserved": 56026,
	"Unassigned": 0,
	"Reserved": 60138,
	"Unassigned": 0,
	"Reserved": 64250,
	"Unassigned": 0,
	"Reserved for Private Use": 65280,
	"renegotiation_info": 65281,
	"Reserved for Private Use": 0,
}
