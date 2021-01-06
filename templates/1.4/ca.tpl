version: '2'
services:
    {{.nodeName}}:
        container_name: {{.nodeName}}
        image: {{.imagePre}}/fabric-ca:{{.imageTag}}
        restart: always
        environment:
          - GODEBUG=netdns=go
          - BCCSP_CRYPTO_TYPE={{.cryptoType}}
          # Issue CN
          - FABRIC_CA_SERVER_CSR_CN={{.nodeName}}
          # Ca root cert expiration
          - FABRIC_CA_SERVER_CSR_CA_EXPIRY=172800h
          # sign enroll cert expiration
          - FABRIC_CA_SERVER_SIGNING_DEFAULT_EXPIRY=172800h
          # sign intermediate cert expiration
          - FABRIC_CA_SERVER_SIGNING_PROFILES_CA_EXPIRY=172800h
           # sign tls cert expiration
          - FABRIC_CA_SERVER_SIGNING_PROFILES_TLS_EXPIRY=172800h
          - FABRIC_CA_SERVER_HOME=/etc/hyperledger/fabric-ca-server
          - FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS=0.0.0.0:9443
        volumes:
          - {{.mountPath}}/{{.nodeName}}:/etc/hyperledger/fabric-ca-server
        command: sh -c 'fabric-ca-server start -b {{.adminName}}:{{.adminPw}}'
        ports:{{range $index,$value:= .ports}}
          - {{$value}}{{end}}
        logging:
          driver: "json-file"
          options:
            max-size: "50m"
            max-file: "5"
        networks:
          - outside

networks:
  outside:
    external:
      name: fabric_network




