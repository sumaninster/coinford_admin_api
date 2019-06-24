bee run -downdoc=true -gendoc=true
openssl genrsa -out app.rsa 2048
openssl rsa -in app.rsa -pubout > app.rsa.pub

13.59.251.209:8080/swagger

-------Beego-------
bee api go-mailcoin-api

-------Ethereum-------
geth --testnet --port=30304 --fast --cache=512 --rpc --rpcport=8547 --rpcapi="db,eth,net,web3,personal,web3"
geth attach http://localhost:8547
personal.newAccount()

Test Ethers:
http://faucet.ropsten.be:3001/

-------Bitcoin-------
https://bitcoin.org/en/full-node#ubuntu-1610
/home/ubuntu/.bitcoin/bitcoin.conf
rpcuser=                                                                                   
rpcpassword=
testnet=1

bitcoind -daemon -testnet
bitcoin-cli help
bitcoin-cli encryptwallet <password>
bitcoin-cli getnewaddress

Test Bitcoins:
https://testnet.coinfaucet.eu/en/

-------Postgres-------
http://suite.opengeo.org/docs/latest/dataadmin/pgGettingStarted/firstconnect.html
sudo -u postgres psql postgres
\password postgres
\q

sudo vi /etc/postgresql/9.6/main/pg_hba.conf
local   	all             all                                md5
host    	all             all        127.0.0.1/32            md5
host    	all             all        0.0.0.0/0               trust
host    	all             all        ::1/128                 md5

/etc/postgresql/9.6/main/postgresql.conf
listen_addresses = '*'

sudo service postgresql restart
psql -U postgres -W
(to test connection)

createdb mailcoin
psql -U postgres mailcoin < path_of_database_dump

-------Firewall-------
sudo firewall-cmd --zone=public --add-port=8090/tcp --permanent
sudo firewall-cmd --reload
iptables-save | grep 8090