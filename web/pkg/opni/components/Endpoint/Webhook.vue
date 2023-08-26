<script>
import { LabeledInput } from '@components/Form/LabeledInput';
import LabeledSelect from '@shell/components/form/LabeledSelect';
import { Checkbox } from '@components/Form/Checkbox';
import BasicAuth from './Auths/BasicAuth';
import Authorization from './Auths/Authorization';
import OAuth2 from './Auths/OAuth2';
import Tls from './Tls';

export default {
  components: {
    Authorization, BasicAuth, Checkbox, LabeledInput, LabeledSelect, OAuth2, Tls
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
        value:   'basicAuth',
        default: { username: '', password: '' }
      },
      {
        label: 'Authorization Header',
        value: 'authorization',
      },
      {
        label:   'OAuth2',
        value:   'oauth2',
        default: { tlsConfig: {} }
      },
    ];

    if (!this.value.httpConfig) {
      this.$set(this.value, 'httpConfig', {
        ...{
          basicAuth:     {},
          authorization: {},
          oauth2:        {},
          tlsConfig:     {}
        },
        ...(this.value.httpConfig || {})
      });
    }

    const authorizationType = this.getAuthType(this.value.httpConfig);

    if (!this.value.httpConfig[authorizationType]) {
      this.$set(this.value.httpConfig, authorizationType, authorizationTypes.find(type => type.value === authorizationType).default);
    }

    return {
      authorizationTypes,
      authorizationType,
    };
  },
  watch: {
    authorizationType(newType, oldType) {
      this.$set(this.value.httpConfig, newType, this.authorizationTypes.find(type => type.value === newType).default || {});
      this.$set(this.value.httpConfig, oldType, {});
    }
  },
  methods: {
    getAuthType(httpConfig) {
      if (httpConfig.authorization?.credentials || httpConfig.authorization?.credentialsFile) {
        return 'authorization';
      }

      if (httpConfig.oauth2?.clientId) {
        return 'oauth2';
      }

      return 'basicAuth';
    },
  }
};
</script>
<template>
  <div>
    <div class="row mt-20 bottom mb-10">
      <div class="col span-6">
        <LabeledInput v-model="value.url" label="URL" :required="true" />
      </div>
      <div class="col span-6">
        <LabeledInput v-model="value.maxAlerts" label="Max Alerts" tooltip="The maximum number of alerts to include in a single webhook message. When the value is 0, all alerts are included." />
      </div>
    </div>
    <h4>Networking</h4>
    <div class="row mt-10 mb-10">
      <div class="col span-6">
        <LabeledInput v-model="value.httpConfig.proxyUrl" label="Proxy URL" />
      </div>
      <div class="col span-2 middle">
        <Checkbox v-model="value.httpConfig.enabledHttp2" label="Use HTTP2" />
      </div>
      <div class="col span-3 middle">
        <Checkbox v-model="value.httpConfig.followRedirects" label="Follow Redirects" />
      </div>
    </div>
    <Tls v-model="value.httpConfig.tlsConfig" class="mb-10" />
    <h4>Authorization</h4>
    <div class="row mt-10">
      <LabeledSelect v-model="authorizationType" :options="authorizationTypes" label="Authorization" />
    </div>
    <BasicAuth v-if="authorizationType === 'basicAuth'" v-model="value.httpConfig.basicAuth" class="mt-10" />
    <Authorization v-if="authorizationType === 'authorization'" v-model="value.httpConfig.authorization" class="mt-10" />
    <OAuth2 v-if="authorizationType === 'oauth2'" v-model="value.httpConfig.oauth2" class="mt-10" />
  </div>
</template>
