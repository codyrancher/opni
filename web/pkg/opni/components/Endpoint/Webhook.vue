<script>
import { LabeledInput } from '@components/Form/LabeledInput';
import LabeledSelect from '@shell/components/form/LabeledSelect';
import { Checkbox } from '@components/Form/Checkbox';

export default {
  components: {
    Checkbox, LabeledInput, LabeledSelect
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
        label: 'Basic Authorization',
        value: 'basic_auth'
      },
      {
        label: 'Authorization Header',
        value: 'authorization'
      },
      {
        label: 'OAuth2',
        value: 'oauth2'
      },
    ];

    const passwordTypes = [
      {
        label: 'Password',
        value: 'password'
      },
      {
        label: 'Password File',
        value: 'password_file'
      },
    ];

    return {
      authorizationTypes,
      authorizationType: authorizationTypes[0].value,
      passwordTypes,
      passwordType:      passwordTypes[0].value
    };
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
    <div class="row mt-10">
      <div class="col span-6">
        <LabeledInput v-model="value.smtpSmartHost" label="Smart Host" :required="true" />
      </div>
      <div class="col span-6">
        <LabeledInput v-model="value.smtpAuthIdentity" label="Identity" :required="true" />
      </div>
    </div>
    <div class="row mt-10">
      <div class="col span-6">
        <LabeledInput v-model="value.smtpAuthUsername" label="Username" :required="true" />
      </div>
      <div class="col span-6">
        <LabeledInput v-model="value.smtpAuthPassword" label="Password" type="password" :required="true" />
      </div>
    </div>
    <div class="row mt-10">
      <div class="col span-6">
        <Checkbox v-model="value.smtpRequireTLS" label="Required TLS" />
      </div>
    </div>
  </div>
</template>
