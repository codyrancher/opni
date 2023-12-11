import { isStandalone } from './utils/standalone';
import MonitoringBackend from './pages/MonitoringBackend';
import Clusters from './pages/Clusters';
import Cluster from './pages/Cluster';

import MetricRoles from './pages/metrics/Roles';
import MetricRole from './pages/metrics/Role';
import MetricRoleBindings from './pages/metrics/RoleBindings';
import MetricRoleBinding from './pages/metrics/RoleBinding';

import Configuration from './pages/Configuration';
import LoggingBackend from './pages/LoggingBackend';
import SLOs from './pages/SLOs';
import SLO from './pages/SLO';
import WorkloadModelConfig from './pages/WorkloadModelConfig';

import Endpoints from './pages/Endpoints';
import Endpoint from './pages/Endpoint';
import Alarms from './pages/Alarms';
import Alarm from './pages/Alarm';
import AlertingOverview from './pages/AlertingOverview';
import AlertingBackend from './pages/AlertingBackend';
import LogAnomalyBackend from './pages/LogAnomalyBackend';
import PretrainedModels from './pages/PretrainedModels';
import Snapshots from './pages/Snapshots';
import Snapshot from './pages/Snapshot';
import Root from './pages/Root';

export const NAVIGATION = {
  routes: [
    {
      name:      isStandalone() ? 'index' : 'opni',
      path:      '/',
      component: Root,
      display:   false
    },
    {
      name:      'agents',
      path:      '/agents',
      labelKey:  'opni.nav.agents',
      component: Clusters,
      routes:    [
        {
          name:      'agent-create',
          path:      '/create',
          labelKey:  'opni.nav.clusters',
          component: Cluster,
          display:   false
        },
      ]
    },
    {
      name:      'logging-config',
      path:      '/logging',
      labelKey:  'opni.nav.loggingConfig',
      component: LoggingBackend,
      display:   true,
      routes:    [
        {
          name:      'snapshots',
          path:      '/snapshots',
          labelKey:  'opni.nav.snapshots',
          component: Snapshots,
          display:   true,
          routes:     [
            {
              name:      'snapshot-create',
              path:      '/create',
              component: Snapshot,
              display:   false,
            },
            {
              name:      'snapshot',
              path:      '/:id',
              component: Snapshot,
              display:   false
            },
          ]
        }
      ]
    },
    {
      name:      'monitoring',
      path:      '/monitoring',
      labelKey:  'opni.nav.monitoring',
      component: MonitoringBackend,
      routes:    [
        {
          name:     'rbac',
          path:     '/rbac',
          labelKey: 'opni.nav.rbac',
          display:   true,
          redirect: { name: 'roles' },
          routes:   [
            {
              name:      'roles',
              path:      '/roles',
              labelKey:  'opni.nav.roles',
              component: MetricRoles,
              routes:    [
                {
                  name:      'role-create',
                  path:      '/create',
                  labelKey:  'opni.nav.roles',
                  component: MetricRole,
                  display:   false
                },
              ]
            },
            {
              name:      'role-bindings',
              path:      '/role-bindings',
              labelKey:  'opni.nav.roleBindings',
              component: MetricRoleBindings,
              routes:    [
                {
                  name:      'role-binding-create',
                  path:      '/create',
                  labelKey:  'opni.nav.roleBindings',
                  component: MetricRoleBinding,
                  display:   false
                },
              ]
            },
          ]
        },
      ]
    },
    {
      name:      'alerting',
      path:      '/alerting',
      labelKey:  'opni.nav.alerting',
      display:   true,
      component: AlertingBackend,
      // redirect:  { name: 'slos' },
      routes:    [
        {
          name:      'alerting-overview',
          path:      '/overview',
          labelKey:  'opni.nav.alertingOverview',
          component: AlertingOverview,
          display:   true
        },
        {
          name:      'endpoints',
          path:      '/endpoints',
          labelKey:  'opni.nav.endpoints',
          component: Endpoints,
          display:   true,
          routes:    [
            {
              name:      'endpoint-create',
              path:      '/create',
              labelKey:  'opni.nav.endpoints',
              component: Endpoint,
              display:   false
            },
            {
              name:      'endpoint',
              path:      '/:id',
              labelKey:  'opni.nav.endpoints',
              component: Endpoint,
              display:   false
            },
          ]
        },
        {
          name:      'alarms',
          path:      '/alarms',
          labelKey:  'opni.nav.alarms',
          component: Alarms,
          display:   true,
          routes:    [
            {
              name:      'alarm-create',
              path:      '/create',
              labelKey:  'opni.nav.alarms',
              component: Alarm,
              display:   false
            },
            {
              name:      'alarm',
              path:      '/:id',
              labelKey:  'opni.nav.alarms',
              component: Alarm,
              display:   false
            },
          ]
        },
        {
          name:      'slos',
          path:      '/slos',
          labelKey:  'opni.nav.slos',
          component: SLOs,
          display:   true,
          routes:    [
            {
              name:      'slo',
              path:      '/:id',
              labelKey:  'opni.nav.slos',
              component: SLO,
              display:   false
            },
            {
              name:      'slo-create',
              path:      '/create',
              labelKey:  'opni.nav.slos',
              component: SLO,
              display:   false
            },
          ]
        },
      ]
    },
    {
      name:      'ai-ops',
      path:      '/ai-ops',
      labelKey:  'opni.nav.aiOps',
      display:   true,
      redirect:  { name: 'log-anomaly' },
      routes:    [
        {
          name:      'log-anomaly',
          path:      '/log-anomaly',
          labelKey:  'opni.nav.logAnomaly',
          display:   true,
          component: LogAnomalyBackend,
          routes:    [
            {
              name:      'pretrained-models',
              path:      '/pretrained-models',
              labelKey:  'opni.nav.pretrainedModels',
              display:   true,
              component: PretrainedModels,
            },
            {
              name:      'workload-model-config',
              path:      '/workload-insights',
              labelKey:  'opni.nav.workloadModel',
              component: WorkloadModelConfig,
              display:   true
            },
          ]
        },

      ]
    },
    {
      name:      'configuration',
      path:      '/configuration',
      labelKey:  'opni.nav.configuration',
      component: Configuration,
      display:   true
    },
  ]
};
