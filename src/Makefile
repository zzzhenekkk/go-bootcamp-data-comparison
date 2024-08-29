.PHONY: all readDB compareDB compareFS

all: readDB compareDB compareFS

readDB:
	@echo "Compiling readDB..."
	@cd cmd/readDB && go build -o ../../readDB
	@echo "readDB compiled successfully!"

compareDB:
	@echo "Compiling compareDB..."
	@cd cmd/compareDB && go build -o ../../compareDB
	@echo "compareDB compiled successfully!"

compareFS:
	@echo "Compiling compareFS..."
	@cd cmd/compareFS && go build -o ../../compareFS
	@echo "compareFS compiled successfully!"

clean:
	@echo "Deleted executable files"
	@rm -rf readDB compareDB compareFS

testReadDB: readDB
	 ./readDB -f testdata/caker.xml
	 ./readDB -f testdata/caker.json

testCompareDB: compareDB
	./compareDB --old testdata/cake.xml --new testdata/cake.json

testCompareFS: compareFS
	./compareFS --old testdata/1.txt --new testdata/2.txt