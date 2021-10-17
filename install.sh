cd /home/$USER/
mkdir go
cd /go
mkdir src
cd /src
mkdir github.com
cd github.com
mkdir Tochka
cd Tochka
cp $PWD/* /
cd ../../..
mkdir github.com
cd github.com
mkdir go-redis
cd go-redis
git clone https://github.com/go-redis/redis
cd ../../..
cd https://github.com/d3417/tochka-source
go build
sudo -u postgres createuser root #enter the current username
sudo -u postgres createdb go_t
sudo -u postgres psql go_t < dumps/cities.sql
sudo -u postgres psql go_t < dumps/countries.sql
cp settings.json.example settings.json
cd /home/$USER/go/src/github.com/d3417/tochka-source
./tochka-free-market sync-models
./tochka-free-market sync-views


