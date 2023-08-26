<script>
import { LabeledInput } from '@components/Form/LabeledInput';
import LabeledSelect from '@shell/components/form/LabeledSelect';
import { Checkbox } from '@components/Form/Checkbox';
import BasicAuth from './Auths/BasicAuth';

export default {
  components: {
    BasicAuth, Checkbox, LabeledInput, LabeledSelect
  },

  props: {
    value: {
      type:     Object,
      required: true
    }
  },

  data() {
    const authorizationTypes = [
      {
        label:   'Basic Authorization',
        value:   'basic_auth',
        default: { username: '', password: '' }
      },
      {
        label: 'Authorization Header',
        value: 'authorization',
      },
      {
        label: 'OAuth2',
        value: 'oauth2'
      },
    ];

    const authorizationType = authorizationTypes[0].value;

    if (!this.value[authorizationType]) {
      this.$set(this.value, authorizationType, authorizationTypes.find(type => type.value === authorizationType).default);
    }

    return {
      authorizationTypes,
      authorizationType,
    };
  },
  watch: {
    authorizationType(newType, oldType) {
      this.$set(this.value, newType, this.authorizationTypes.find(type => type.value === newType).default);
      this.$set(this.value, oldType, {});
    }
  }
};
</script>
<template>
  <div>
    <div class="row mt-20 bottom mb-10">
      <div class="col span-6">
        <LabeledInput v-model="value.url" label="Url" :required="true" />
      </div>
      <div class="col span-6">
        <LabeledInput v-model="value.max_alerts" label="Max Alerts" />
      </div>
    </div>
    <h4>Authorization</h4>
    <div class="row mt-10">
      <LabeledSelect v-model="authorizationType" :options="authorizationTypes" label="Authorization" />
    </div>
    <BasicAuth v-if="authorizationType === 'basic_auth'" v-model="value.basic_auth" class="mt-10" />
  </div>
</template>
