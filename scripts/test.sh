OUTPUT="Package ID: basic_1.0:69de748301770f6ef64b42aa6bb6cb291df20aa39542c3ef94008615704007f3, Label: basic_1.0"
PACKAGE_ID=$(sed -n 's/.*Package ID: \(.*\), Label:.*/\1/p' <<< "$OUTPUT")

echo $PACKAGE_ID