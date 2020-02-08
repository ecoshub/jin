const fs = require('fs');

let jsonFile  = 'test/test-json.json';
let pathFile  = 'test/test-json-paths.json';
let valueFile = 'test/test-json-values.json';
let pathString = '';
let valueString = '';

const pathToString = (arr) => {
	var result = '[';
	arr.forEach((e) => {result += JSON.stringify(e) + ','});
	if (result.length > 1) {
		result = result.slice(0, result.length - 1);
	}
	result += ']';
	return result;
}

const getType = (json) => {
	if (typeof json === 'object' ){
		if (Array.isArray(json)){
			return 'array';
		}
		return 'object';
	}
	return 'value';
}

// core walk func to walk through all elements and get specific properties.
const walk = (json, paths, upType, myType, call) => {
	// getting keys and values.
	const keys = Object.keys(json);
	const values = Object.values(json);
	// upType is defines what this element in.
	let upTypeIs = getType(json);
	for (var i = 0; i < keys.length; i++) {
		let val = values[i];
		let key = keys[i];
		// elements type
		let myTypeIs = getType(val);
		// for path creation push last key and pop after callback.
		paths.push(key);
		call(key, val, upTypeIs, myTypeIs, paths);
		paths.pop();
		// if it is an iterable object than recursively walk through.
		if (typeof val === 'object' && val != null && val != undefined) {
			// standart path creation.
			paths.push(key);
			walk(val, paths, upTypeIs, myTypeIs, call);
			paths.pop();
		}
	}
}

const addPathAndValue = (path, value) => {
	pathString += pathToString(path) + '\n';
	valueString += JSON.stringify(value) + '\n';
}

const createCase = (json, pathType) => {
	switch (pathType) {
	case "all":
		addPathAndValue([], json)
		walk(json, [], getType(json), null, (key, value, upType, myType, path) => {
			addPathAndValue(path, value)
		});
		break;
	case "array":
		if (getType(json) === 'array'){
			addPathAndValue([], json)
		}
		walk(json, [], getType(json), null, (key, value, upType, myType, path) => {
			if (myType === 'array'){
				addPathAndValue(path, value)
			}
		});
		break;
	case "object":
		if (getType(json) === 'object'){
			addPathAndValue([], json)
		}
		walk(json, [], getType(json), null, (key, value, upType, myType, path) => {
			if (myType === 'object'){
				addPathAndValue(path, value)
			}
		});
		break;
	case "object-values":
		walk(json, [], getType(json), null, (key, value, upType, myType, path) => {
			if (upType === 'object'){
				addPathAndValue(path, value)
			}
		});
		break;
	case "array-values":
		walk(json, [], getType(json), null, (key, value, upType, myType, path) => {
			if (upType === 'array'){
				addPathAndValue(path, value)
			}
		});
		break;
	case "not-value":
		walk(json, [], getType(json), null, (key, value, upType, myType, path) => {
			if (myType !== 'value'){
				addPathAndValue(path, value)
			}
		});
		break;
	}
	if (pathString.length == 0) {
		fs.writeFileSync(pathFile, "");
		fs.writeFileSync(valueFile, "");
	} else {
		fs.writeFileSync(pathFile, pathString.slice(0, pathString.length - 1));
		fs.writeFileSync(valueFile, valueString.slice(0, valueString.length - 1));
	}
}

if (process.argv.length > 2) {
	var json = JSON.parse(fs.readFileSync(jsonFile));
	createCase(json, process.argv[2]);
}