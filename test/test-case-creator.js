const fs = require('fs');
var exec = require('child_process').execFile;

let jsonFile  = 'test/test-json.json';
let pathFile  = 'test/test-json-paths.json';
let valueFile = 'test/test-json-values.json';

const pathToString = (arr) => {
	var result = "[";
	arr.forEach((e) => {
		result += JSON.stringify(e) + ",";
	})
	result = result.slice(0, result.length - 1);
	result += "]";
	return result;
}

const createPaths = (obj, key, tag) => {
	const keys = Object.keys(obj);
	const values = Object.values(obj);
	for ( var i = 0 ; i < keys.length ; i ++ ){
		if (tag === 'all') {
			key.push(keys[i]);
			let path = pathToString(key);
			fs.appendFileSync(pathFile, path + `\n`);
			key.pop();
		}
		if (tag === 'objects') {
			if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined && !Array.isArray(values[i])) {
				key.push(keys[i]);
				let path = pathToString(key);
				fs.appendFileSync(pathFile, path + `\n`);
				key.pop();
			}
		}
		if (tag === 'arrays') {
			if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined && Array.isArray(values[i])) {			
				key.push(keys[i]);
				let path = pathToString(key);
				fs.appendFileSync(pathFile, path + `\n`);
				key.pop();
			}
		}
		if (tag === 'values') {
			if (typeof values[i] != 'object' && values[i] != null && values[i] != undefined) {
				key.push(keys[i]);
				let path = pathToString(key);
				fs.appendFileSync(pathFile, path + `\n`);
				key.pop();
			}
		}
		if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined) {
			key.push(keys[i]);
			createPaths(values[i], key, tag);
			key.pop();
		}
	}
}

const createValues = (obj, tag, func) => {
	const keys = Object.keys(obj)
	const values = Object.values(obj)
	for ( var i = 0 ; i < keys.length ; i ++ ){
		if (tag === 'all') {
			values[i] = func(values[i]);
			fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
		}
		if (tag === 'objects') {
			if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined && !Array.isArray(values[i])) {
				values[i] = func(values[i]);
				fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
			}
		}
		if (tag === 'arrays') {
			if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined && Array.isArray(values[i])) {
				values[i] = func(values[i]);
				fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
			}
		}
		if (tag === 'values') {
			if (typeof values[i] != 'object' && values[i] != null && values[i] != undefined) {
				values[i] = func(values[i]);
				fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
			}
		}
		if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined) {
			createValues(values[i], tag, func);
		}
	}
}

function addKeyValue(obj, key, val) {
	const keys = Object.keys(obj);
	const values = Object.values(obj);
	for ( var i = 0 ; i < keys.length ; i ++ ){
		if (typeof values[i] === 'object' && values[i] != null && values[i] != undefined && !Array.isArray(values[i])) {
			values[i][key] = val;
			fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
			addKeyValue(values[i], key, val);
		}
	}
	return obj;
}

const clearNewLine = (dir) => {
	let file  = fs.readFileSync(dir);
	file = file.slice(0, file.length - 1);
	fs.writeFileSync(dir, file.toString());
}

const createTestCase = (json, content, func) => {
	fs.writeFileSync(valueFile, "");
	fs.writeFileSync(pathFile, "");
	createValues(json, content, func);
	createPaths(json, [], content);
	clearNewLine(valueFile);
	clearNewLine(pathFile);
}

if (process.argv.length === 3) {
	var mainArray = JSON.parse(fs.readFileSync(jsonFile));
	if (process.argv[2] === "get"){
		createTestCase(mainArray, 'all', (val)=>{return val})
	}
	if (process.argv[2] === "set"){
		createTestCase(mainArray, 'values', (val) => {
			return 'test-string';
		})
	}
}

