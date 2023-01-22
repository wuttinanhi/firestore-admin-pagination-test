run:
	export GOOGLE_APPLICATION_CREDENTIALS=C:/secrets/serviceAccountKey.json && go run .

build:
	export GOOGLE_APPLICATION_CREDENTIALS=C:/secrets/serviceAccountKey.json && go build
