package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gaukas/tlsspecloader"
)

var (
	parseClientCertificateType = flag.Bool("clientcerttype", false, "Parse ClientCertificateType Identifiers")
	parseCipherSuite           = flag.Bool("ciphersuite", false, "Parse CipherSuite Identifiers")
	parseContentType           = flag.Bool("contenttype", false, "Parse ContentType")
	parseAlert                 = flag.Bool("alert", false, "Parse Alerts")
	parseHandshakeType         = flag.Bool("hstype", false, "Parse HandshakeType")
	parseSupportedGroups       = flag.Bool("supportedgroups", false, "Parse Supported Groups")
	parseEcPoint               = flag.Bool("ecpoint", false, "Parse EC Point Formats")
	parseEcCurve               = flag.Bool("eccurve", false, "Parse EC Curve Types")
	parseSupplementalData      = flag.Bool("suppdata", false, "Parse Supplemental Data Formats")
	parseUserMappingType       = flag.Bool("usermapping", false, "Parse User Mapping Type Values")
	parseSignatureAlgorithm    = flag.Bool("sigalg", false, "Parse SignatureAlgorithm")
	parseHashAlgorithm         = flag.Bool("hashalg", false, "Parse HashAlgorithm")
	parseAuthorizationData     = flag.Bool("authdata", false, "Parse Authorization Data Formats")
	parseHeartBeatMessage      = flag.Bool("hbmsg", false, "Parse HeartBeat Message Types")
	parseHeartBeatMode         = flag.Bool("hbmode", false, "Parse HeartBeat Modes")
	parseSignatureScheme       = flag.Bool("sigscheme", false, "Parse SignatureScheme")

	parseExtensionType = flag.Bool("exttype", false, "Parse ExtensionType Values")
)

func main() {
	flag.Parse()

	if *parseClientCertificateType {
		// ClientCertificateType Identifiers
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-2.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/clientcerttype.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-2.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseCipherSuite {
		// CipherSuite Identifiers
		// "0x00,0x07" -> 0x0007 (uint16 hex, uint16 hex)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-4.csv")
		if err != nil {
			panic(err)
		}

		hex := tlsspecloader.HexFormatter{}
		ids := hex.DuoToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{HexOutput: true}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/ciphersuite.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-4.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseContentType {
		// ContentType
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-5.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/contenttype.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-5.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseAlert {
		// Alert
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-6.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/alert.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-6.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseHandshakeType {
		// HandshakeType
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-7.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/handshaketype.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-7.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseSupportedGroups {
		// SupportedGroups
		// "6550" -> 6550 (uint16 dec, uint16 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-8.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/supported_groups.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-8.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseEcPoint {
		// EcPoint
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-9.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/ec_point_formats.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-9.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseEcCurve {
		// EcCurve
		// "123" -> 123 (uint16 dec, uint16 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-10.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/ec_curve_types.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-10.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseSupplementalData {
		// SupplementalDataType/Supplemental Data Formats
		// "1230" -> 1230 (uint16 dec, uint16 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-12.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/supplemental_data_formats.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-12.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseUserMappingType {
		// UserMappingType Values
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-14.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/usermappingtype_values.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-14.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseSignatureAlgorithm {
		// SignatureAlgorithm
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-16.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/signaturealgorithm.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-16.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseHashAlgorithm {
		// HashAlgorithm
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("tls-parameters-18.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/hashalgorithm.go",
			"https://www.iana.org/assignments/tls-parameters/tls-parameters-18.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseAuthorizationData {
		// AuthorizationDataFormat
		// "123" -> 123 (uint16 dec, uint16 dec)
		csv, err := tlsspecloader.ReadCSV("authorization-data.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/authorization_data_formats.go",
			"https://www.iana.org/assignments/tls-parameters/authorization-data.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseHeartBeatMessage {
		// HeartBeatMessage
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("heartbeat-message-types.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/heartbeat_message_types.go",
			"https://www.iana.org/assignments/tls-parameters/heartbeat-message-types.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseHeartBeatMode {
		// HeartBeatMode
		// "123" -> 123 (uint8 dec, uint8 dec)
		csv, err := tlsspecloader.ReadCSV("heartbeat-modes.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint8Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint8(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint8(ids, csv.Column(1))

		generateFile(
			"out/heartbeat_mode.go",
			"https://www.iana.org/assignments/tls-parameters/heartbeat-modes.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseSignatureScheme {
		// SignatureScheme
		// "123" -> 123 (uint16 dec, uint16 dec)
		csv, err := tlsspecloader.ReadCSV("tls-signaturescheme.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/signaturescheme.go",
			"https://www.iana.org/assignments/tls-parameters/tls-signaturescheme.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}

	if *parseExtensionType {
		// ExtensionType
		// "123" -> 123 (uint16 dec, uint16 dec)
		csv, err := tlsspecloader.ReadCSV("tls-extensiontype-values-1.csv")
		if err != nil {
			panic(err)
		}

		dec := tlsspecloader.DecFormatter{}
		ids := dec.ToUint16Arr(csv.Column(0))
		loader := tlsspecloader.Go{}
		enum := loader.EnumUint16(ids, csv.Column(1))
		idIdxName, nameIdxID := loader.MutualMapUint16(ids, csv.Column(1))

		generateFile(
			"out/extensiontype.go",
			"https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values-1.csv",
			enum,
			idIdxName,
			nameIdxID,
		)
	}
}

func generateFile(
	filename string,
	source string,
	enum string,
	idIdxName string,
	nameIdxID string,
) {
	// Write to file
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	f.WriteString("package godicttls\n\n")
	f.WriteString(fmt.Sprintf("// source: %s\n", source))
	f.WriteString(fmt.Sprintf(
		"// last updated: %s\n\n",
		time.Now().Format("January 2006"),
	)) // last updated: March 2023
	f.WriteString(enum)
	f.WriteString("\n")
	f.WriteString(idIdxName)
	f.WriteString("\n")
	f.WriteString(nameIdxID)
	f.Close()
}
