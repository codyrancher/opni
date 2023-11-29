<script>
import Main from '@pkg/opni/components/Main';
import { choosePageLayout } from '@pkg/opni/utils/standalone';
import RoleBinding from '@pkg/opni/components/RoleBinding';

export default {
  components: { RoleBinding, Main },
  layout:     choosePageLayout(),
  head() {
    return { title: 'Opni - Role Binding' };
  }
};
</script>
<template>
  <Main>
    <RoleBinding backend="metrics" />
  </Main>
</template>
