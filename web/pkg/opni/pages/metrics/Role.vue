<script>
import Main from '@pkg/opni/components/Main';
import Role from '@pkg/opni/components/Role';
import { choosePageLayout } from '@pkg/opni/utils/standalone';

export default {
  components: { Role, Main },
  layout:     choosePageLayout(),
  head() {
    return { title: 'Opni - Role' };
  }
};
</script>
<template>
  <Main>
    <Role backend="metrics" />
  </Main>
</template>
