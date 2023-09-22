// @generated by service-generator v0.0.1 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/plugins/metrics/apis/cortexops/cortexops.proto (package cortexops, syntax proto3)
/* eslint-disable */

import { ConfigurationHistoryRequest, GetRequest, InstallStatus } from "../../../../pkg/plugins/driverutil/types_pb";
import { CapabilityBackendConfigSpec, ConfigurationHistoryResponse, DryRunRequest, DryRunResponse, PresetList, ResetRequest } from "./cortexops_pb";
import { axios } from "@pkg/opni/utils/axios";


export async function GetDefaultConfiguration(input: GetRequest): Promise<CapabilityBackendConfigSpec> {
  try {
    return (await axios.request({
    transformResponse: resp => CapabilityBackendConfigSpec.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration/default`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function SetDefaultConfiguration(input: CapabilityBackendConfigSpec): Promise<void> {
  try {
    return (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration/default`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function ResetDefaultConfiguration(): Promise<void> {
  try {
    return (await axios.request({
      method: 'delete',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration/default`
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetConfiguration(input: GetRequest): Promise<CapabilityBackendConfigSpec> {
  try {
    return (await axios.request({
    transformResponse: resp => CapabilityBackendConfigSpec.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function SetConfiguration(input: CapabilityBackendConfigSpec): Promise<void> {
  try {
    return (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function ResetConfiguration(input: ResetRequest): Promise<void> {
  try {
    return (await axios.request({
      method: 'delete',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function Status(): Promise<InstallStatus> {
  try {
    return (await axios.request({
    transformResponse: resp => InstallStatus.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/status`
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function Install(): Promise<void> {
  try {
    return (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/install`
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function Uninstall(): Promise<void> {
  try {
    return (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/uninstall`
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function ListPresets(): Promise<PresetList> {
  try {
    return (await axios.request({
    transformResponse: resp => PresetList.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/presets`
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function DryRun(input: DryRunRequest): Promise<DryRunResponse> {
  try {
    return (await axios.request({
    transformResponse: resp => DryRunResponse.fromBinary(new Uint8Array(resp)),
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/dry-run`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}


export async function ConfigurationHistory(input: ConfigurationHistoryRequest): Promise<ConfigurationHistoryResponse> {
  try {
    return (await axios.request({
    transformResponse: resp => ConfigurationHistoryResponse.fromBinary(new Uint8Array(resp)),
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/CortexOps/configuration/history`,
    data: input?.toBinary() as ArrayBuffer
    })).data;
  } catch (ex) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, new Uint8Array(ex?.response?.data));
      console.error(s);
    }
    throw ex;
  }
}
