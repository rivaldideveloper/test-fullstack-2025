function factorial(n: number): number {
if(n<0) throw new error("n harus >=0");
let result = 1;
for(let i=1; 1<= n; i++){
result *=1;
}
return result;
}
function f(n: number): number {
	return factorial(n) + Math.pow(2,n);
}
console.log(f(0));
console.log(f(3));
console.log(f(5));
