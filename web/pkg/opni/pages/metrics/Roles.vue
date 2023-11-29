<script>
import Main from '@pkg/opni/components/Main';
import { choosePageLayout } from '@pkg/opni/utils/standalone';
import Roles from '@pkg/opni/components/Roles';

export default {
  components: { Roles, Main },
  layout:     choosePageLayout(),
  head() {
    return { title: 'Opni - Roles' };
  }
};
</script>
<template>
  <Main>
    <Roles backend="metrics" />
  </Main>
</template>
