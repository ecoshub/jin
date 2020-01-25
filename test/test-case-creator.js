const fs = require('fs');
var exec = require('child_process').execFile;

let jsonFile  = 'test/test-json.json';
let pathFile  = 'test/test-json-paths.json';
let valueFile = 'test/test-json-values.json';

const pathToString = (arr) => {
	var result = '[';
	arr.forEach((e) => {result += JSON.stringify(e) + ',';})
	result = result.slice(0, result.length - 1);
	result += ']';
	return result;
}

const createPaths = (obj, type, key, tag, func) => {
	let lastType = type;
	const keys = Object.keys(obj);
	const values = Object.values(obj);
	for ( var i = 0 ; i < keys.length ; i ++ ){
		if (values[i] !== null && values[i] !== undefined){
			if (tag === 'all') {
				key.push(keys[i]);
				let path = pathToString(key);
				fs.appendFileSync(pathFile, path + `\n`);
				values[i] = func(values[i]);
				if ( values[i] !== undefined){
					fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
				}
				key.pop();
			}
			if (tag === 'object') {
				if (typeof values[i] === 'object' && !Array.isArray(values[i])) {
					key.push(keys[i]);
					let path = pathToString(key);
					fs.appendFileSync(pathFile, path + `\n`);
					values[i] = func(values[i]);
					if ( values[i] !== undefined){
						fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
					}
					key.pop();
				}
			}
			if (tag === 'array') {
				if (typeof values[i] === 'object' && Array.isArray(values[i])) {			
					key.push(keys[i]);
					let path = pathToString(key);
					fs.appendFileSync(pathFile, path + `\n`);
					values[i] = func(values[i]);
					if ( values[i] !== undefined){
						fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
					}
					key.pop();
				}
			}
			if (tag === 'arrayvalues') {
				if (type === 'array'){
						key.push(keys[i]);
						let path = pathToString(key);
						fs.appendFileSync(pathFile, path + `\n`);
						values[i] = func(values[i]);
						if ( values[i] !== undefined){
							fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
						}
						key.pop();
				}
			}
			if (tag === 'objectvalues') {
				if (type === 'object'){
						key.push(keys[i]);
						let path = pathToString(key);
						fs.appendFileSync(pathFile, path + `\n`);
						values[i] = func(values[i]);
						if ( values[i] !== undefined){
							fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
						}
						key.pop();
				}
			}
			if (tag === 'values') {
				if (typeof values[i] !== 'object') {
					key.push(keys[i]);
					let path = pathToString(key);
					fs.appendFileSync(pathFile, path + `\n`);
					values[i] = func(values[i]);
					if ( values[i] !== undefined){
						fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
					}
					key.pop();
				}
			}
			if (tag === 'notvalues') {
				if (typeof values[i] === 'object') {
					key.push(keys[i]);
					let path = pathToString(key);
					fs.appendFileSync(pathFile, path + `\n`);
					values[i] = func(values[i]);
					if ( values[i] !== undefined){
						fs.appendFileSync(valueFile, JSON.stringify(values[i]) + `\n`);
					}
					key.pop();
				}
			}
			if (typeof values[i] === 'object') {
				if (Array.isArray(values[i])){
					type = 'array';
				}else{
					type = 'object';
				}
				key.push(keys[i]);
				createPaths(values[i], type, key, tag, func);
				key.pop();
				type = lastType;
			}
		}
	}
}

const clearNewLine = (dir) => {
	let file  = fs.readFileSync(dir);
	file = file.slice(0, file.length - 1);
	fs.writeFileSync(dir, file.toString());
}

const createTestCase = (json, content, func) => {
	fs.writeFileSync(pathFile, '');
	fs.writeFileSync(valueFile, '');
	let type = '';
	if (typeof json === 'object'){
		if (Array.isArray(json)){
			type = 'array'
		}else{
			type = 'object'
		}
	}else{
		type = 'value'
	}
	if (content === type || (content === 'arrayvalues' && type === 'array') || (content === 'objectvalues' && type === 'object')) {
		fs.appendFileSync(pathFile, `[]` + '\n');
		fs.appendFileSync(valueFile, JSON.stringify(func(json)) + `\n`);
	}
	createPaths(json, type, [], content, func);
	clearNewLine(pathFile);
	clearNewLine(valueFile);
}

if (process.argv.length > 2) {
	var mainArray = JSON.parse(fs.readFileSync(jsonFile));
	if (process.argv[2] === 'get'){
		createTestCase(mainArray, 'all', (val)=>{return val})
	}
	if (process.argv[2] === 'set'){
		createTestCase(mainArray, 'values', (val) => {
			return 'test-string';
		})
	}
	if (process.argv[2] === 'addkv'){
		createTestCase(mainArray, 'object', (val) => {
			val['test-key'] = 'test-value';
			return val;
		})
	}
	if (process.argv[2] === 'add'){
		createTestCase(mainArray, 'array', (arr) => {
			arr.push('test-value');
			return arr;
		})
	}
	if (process.argv[2] === 'insert'){
		createTestCase(mainArray, 'array', (arr) => {
			arr.splice(0, 0, 'test-value')
			return arr;
		})
	}
	if (process.argv[2] === 'deleteKV'){
		createTestCase(mainArray, 'object', (arr) => {
			return arr;
		})
	}
	if (process.argv[2] === 'deleteV'){
		createTestCase(mainArray, 'array', (arr) => {
			return arr;
		})
	}
	if (process.argv[2] === 'arrayiter'){
		createTestCase(mainArray, 'array', (arr) => {
			return arr;
		})
	}
	if (process.argv[2] === 'objectiter'){
		createTestCase(mainArray, 'object', (arr) => {
			return arr;
		})
	}
}
