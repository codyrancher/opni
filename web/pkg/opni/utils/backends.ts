export const BACKENDS = {
  METRICS: 'metrics',
  LOGGING: 'logging'
};

export function propValidator(value: string) {
  return Object.values(BACKENDS).includes(value);
}
