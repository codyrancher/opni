<script>
import { Checkbox } from '@components/Form/Checkbox';
import { LabeledInput } from '@components/Form/LabeledInput';
import LabeledSelect from '@shell/components/form/LabeledSelect';

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
    if (!this.value.minVersion) {
      this.$set(this.value, 'minVersion', 'TLS12');
    }

    if (!this.value.maxVersion) {
      this.$set(this.value, 'maxVersion', 'TLS13');
    }

    return {
      tlsOptions: [
        {
          label: 'TLS 1.0',
          value: 'TLS10'
        },
        {
          label: 'TLS 1.1',
          value: 'TLS11'
        },
        {
          label: 'TLS 1.2',
          value: 'TLS12'
        },
        {
          label: 'TLS 1.3',
          value: 'TLS13'
        },
      ]
    };
  }
};
</script>
<template>
  <div>
    <h4>TLS</h4>
    <div class="row">
      <div class="col span-6">
        <LabeledSelect v-model="value.minVersion" :options="tlsOptions" label="TLS Min Version" />
      </div>
      <div class="col span-6">
        <LabeledSelect v-model="value.maxVersion" :options="tlsOptions" label="TLS Max Version" />
      </div>
    </div>
    <div class="row mt-10">
      <div class="col span-6">
        <LabeledInput v-model="value.serverName" label="Server Name" />
      </div>
      <div class="col span-6 middle">
        <Checkbox v-model="value.insecureSkipVerify" label="Skip Server Cert Validation" />
      </div>
    </div>
    <div class="row mt-10">
      <div class="col span-4">
        <LabeledInput v-model="value.caFile" label="CA File" />
      </div>
      <div class="col span-4">
        <LabeledInput v-model="value.certFile" label="Cert File" />
      </div>
      <div class="col span-4">
        <LabeledInput v-model="value.keyFile" label="Key File" />
      </div>
    </div>
  </div>
</template>
