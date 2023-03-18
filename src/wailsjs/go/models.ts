export namespace IPC {
	
	export class Request {
	    module: string;
	    method: string;
	    data: {[key: string]: any};
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.module = source["module"];
	        this.method = source["method"];
	        this.data = source["data"];
	    }
	}
	export class response {
	    module: string;
	    method: string;
	    data: {[key: string]: any};
	
	    static createFrom(source: any = {}) {
	        return new response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.module = source["module"];
	        this.method = source["method"];
	        this.data = source["data"];
	    }
	}

}

