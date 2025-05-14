import { ref, computed } from 'vue';
import { regexes } from '~/consts/regex';

export function usePasswordStrength(password: Ref<string | undefined>) {

  function checkStrength(str: string) {
    const requirements = [
      { label: 'minimal', regex: /./, score: 0.1 },
      { label: 'weak', regex: regexes.password.weak, score: 0.2 },
      { label: 'medium', regex: regexes.password.medium, score: 0.6 },
      { label: 'strong', regex: regexes.password.strong, score: 1 },
    ];

    return requirements.map((req) => ({ ...req, met: req.regex.test(str) }));
  }

  const strength = computed(() => checkStrength(password.value || ''));
  const strengthMatched = computed(() => [...strength.value].reverse().find(req => req.met))
  const score = computed(() => strengthMatched.value?.score ?? 0)

  const strengthLabel = computed(() => strengthMatched.value?.label || 'minimal')
  const strengthLevelSwitch = (target: string, values: Array<string>) => {
    switch (target) {
      case 'minimal':
      case 'weak': return values[1]
      case 'medium': return values[2]
      case 'strong': return values[3]
      default: return values[0]
    }
  }

  const color = computed(() => {
    const res = strengthLevelSwitch(strengthLabel.value, ["neutral", "error", "warning", "success"])
    return res as "neutral" | "error" | "warning" | "success"
  })
  const text = computed(() => strengthLevelSwitch(strengthLabel.value, ["Enter a password", "Weak password", "Medium password", "Strong password"]))

  return { score, color, text };
}
