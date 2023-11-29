import { Core } from '@pkg/opni/api/opni';
import { RoleBinding } from './RoleBinding';

export class RoleBindingList {
    base: Core.Types.RoleBindingList;

    constructor(base: Core.Types.RoleBindingList) {
      this.base = base;
    }

    get items(): RoleBinding[] {
      return this.base.items.map(s => new RoleBinding(s, null));
    }
}
