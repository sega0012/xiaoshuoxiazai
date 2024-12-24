export namespace main {
	
	export class Xiaoshuo {
	
	
	    static createFrom(source: any = {}) {
	        return new Xiaoshuo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class XzInfo {
	    name: string;
	    firstUrl: string;
	    headUrl: string;
	    content: string;
	    title: string;
	    nextPage: string;
	    breakFlag: string;
	
	    static createFrom(source: any = {}) {
	        return new XzInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.firstUrl = source["firstUrl"];
	        this.headUrl = source["headUrl"];
	        this.content = source["content"];
	        this.title = source["title"];
	        this.nextPage = source["nextPage"];
	        this.breakFlag = source["breakFlag"];
	    }
	}

}

