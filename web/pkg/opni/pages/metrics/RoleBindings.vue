<script>
import Main from '@pkg/opni/components/Main';
import { choosePageLayout } from '@pkg/opni/utils/standalone';
import RoleBindings from '@pkg/opni/components/RoleBindings';

export default {
  components: { RoleBindings, Main },
  layout:     choosePageLayout(),
  head() {
    return { title: 'Opni - Role Bindings' };
  }
};
</script>
<template>
  <Main>
    <RoleBindings backend="metrics" />
  </Main>
</template>
