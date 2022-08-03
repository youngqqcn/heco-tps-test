
init:
	 cd node1 && geth --datadir=data init ../myheco.json  && cd ..
	 cd node2 && geth --datadir=data init ../myheco.json  && cd ..

start-node1:
	cd node1 && geth --datadir data --networkid 46614 --bootnodes enode://ccfae4d05ad0625d4be99841b7a68caf52020660098dd69ceb5a150e5cb38e0627f8eb4c17e3fbba260856e871851a4305b8f8518d7ea0a821cd149f45b3d1f2@127.0.0.1:0?discport=30301 --ipcdisable --syncmode full --http --allow-insecure-unlock --http.addr 0.0.0.0 --unlock 0x923c36fb1ddd9305128562996c52af951fd469fa --password password.txt --mine console


start-node2:
	cd node2 && geth --datadir data --networkid 46614 --bootnodes enode://ccfae4d05ad0625d4be99841b7a68caf52020660098dd69ceb5a150e5cb38e0627f8eb4c17e3fbba260856e871851a4305b8f8518d7ea0a821cd149f45b3d1f2@127.0.0.1:0?discport=30301 --port 30304 --ipcdisable --syncmode full --http --allow-insecure-unlock --http.addr 0.0.0.0 --http.port 8546  --unlock 0xeb8ac1b19b45367922aae6d5e471ba7ce8f6d15c --password password.txt --mine console

bootnode:
	bootnode -nodekey ./bnode/boot.key

clean:
	rm -rf node1/data/geth && rm node1/data/history && rm -rf node2/data/geth && rm node2/data/history

