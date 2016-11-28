run:
	./runGame.sh
zip:
	zip -r submission`date '+%Y-%m-%d_%H-%M'`.zip MyBot.go src
clean:
	rm -rf *.hlt *.zip
