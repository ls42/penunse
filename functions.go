package main

import (
	"encoding/binary"
	"flag"
	"log"
)

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func parseFlags() params {
	p := params{}
	flag.IntVar(&p.port, "port", 4202, "port to listen on")
	flag.StringVar(&p.dbhost, "dbhost", "localhost", "hostname of database server")
	flag.IntVar(&p.dbport, "dbport", 5432, "tcp port of database server")
	flag.StringVar(&p.dbuser, "dbuser", "penunse", "username for database server")
	flag.StringVar(&p.dbname, "dbname", "penunse", "database name")
	flag.StringVar(&p.dbpass, "dbpass", "foo", "password for database server")
	flag.Parse()
	if err := p.validate(); err != nil {
		log.Fatal(err)
	}
	return p
}
