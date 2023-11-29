import { deleteRoleBinding } from '@pkg/opni/utils/requests/management';
import { Resource } from '@pkg/opni/models/Resource';
import { Core } from '@pkg/opni/api/opni';

export class RoleBinding extends Resource {
    private base: Core.Types.RoleBinding;

    constructor(base: Core.Types.RoleBinding, vue: any) {
      super(vue);
      this.base = base;
    }

    get name() {
      return this.base.id;
    }

    get nameDisplay(): string {
      return this.name;
    }

    get subjects() {
      return this.base.subjects;
    }

    get role() {
      return this.base.roleId;
    }

    get taints() {
      return this.base.taints;
    }

    get availableActions(): any[] {
      return [
        {
          action:     'promptRemove',
          altAction:  'delete',
          label:      'Delete',
          icon:       'icon icon-trash',
          bulkable:   true,
          enabled:    true,
          bulkAction: 'promptRemove',
          weight:     -10, // Delete always goes last
        }
      ];
    }

    async remove() {
      await deleteRoleBinding(this.base.id);
      super.remove();
    }
}
