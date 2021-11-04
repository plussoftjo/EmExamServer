/* Do not change, this code is generated from Golang structs */


export class Answers {
    title: string;
    questions_id: number;
    correct: boolean;

    static createFrom(source: any = {}) {
        return new Answers(source);
    }

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.title = source["title"];
        this.questions_id = source["questions_id"];
        this.correct = source["correct"];
    }
}