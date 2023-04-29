import { Message } from "../format/base";

export class Response {
    success: boolean = false
    message?: Message
    constructor(success: boolean, message?: Message) {
        this.success = success;
        this.message = message;
    }
}