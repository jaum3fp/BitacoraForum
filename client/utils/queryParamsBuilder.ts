export default function buildQueryParams(filter: object) {
  const params = Object.entries(filter)
      .filter(([_, value]) => value !== undefined && value !== null && value !== '')
      .map(([key, value]) => {
        if (Array.isArray(value)) {
          return `${encodeURIComponent(key)}=${encodeURIComponent(value.join(','))}`;
        }
        return `${encodeURIComponent(key)}=${encodeURIComponent(value)}`;
      });

    return params.length ? `?${params.join('&')}` : '';
}
