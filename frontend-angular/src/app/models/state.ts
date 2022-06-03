import {Shipping} from "./shipping";
import {User} from "./user";

export class State {

    id?: string;
    processDefinitionId?: string;
    name?: string;
    activityId?: string;
    activityName?: string;
    childActivityInstances?: State[]
    incidents?: [];
    state?: string;

    constructor() {
    }
}