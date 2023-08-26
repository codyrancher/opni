<script>
import { LabeledInput } from '@components/Form/LabeledInput';
import LabeledSelect from '@shell/components/form/LabeledSelect';

export default {
  components: { LabeledInput, LabeledSelect },

  props: {
    value: {
      type:     Object,
      required: true
    }
  },

  data() {
    const passwordTypes = [
      {
        label: 'Password',
        value: 'password'
      },
      {
        label: 'Password File',
        value: 'password_file'
      }
    ];

    const passwordType = passwordTypes.map(pt => pt.value).find(type => this.value[type]) || passwordTypes[0].value;

    return {
      passwordTypes,
      passwordType
    };
  }
};
</script>
<template>
  <div>
    <div class="row">
      <div class="col span-12">
        <LabeledInput v-model="value.username" label="Username" :required="true" />
      </div>
    </div>
    <div class="row mt-10">
      <div class="col span-4">
        <LabeledSelect v-model="passwordType" :options="passwordTypes" label="Password Type" :required="true" />
      </div>
      <div class="col span-8">
        <LabeledInput v-if="passwordType === 'password'" v-model="value.password" label="Password" :required="true" :password="true" />
        <LabeledInput v-else v-model="value.password_file" label="Password File" :required="true" :password="true" />
      </div>
    </div>
  </div>
</template>
