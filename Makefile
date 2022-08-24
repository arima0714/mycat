build:
	echo "go build を実行し、catを作成する"
	go build main.go
	mv main cat

clean:; rm -rf main cat
