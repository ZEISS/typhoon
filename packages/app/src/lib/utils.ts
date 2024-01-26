import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export const teamsWorkloadUrlPath = '/teams/[team]/workloads/[id]'
export const teamsWorkloadLensPath =
  '/teams/[team]/workloads/[id]/lenses/[lensId]'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function removeTrailingSlash(path: string) {
  return path.endsWith('/') && path.length > 1 ? path.slice(0, -1) : path
}

export function getRouteAsPath(
  pathname: string,
  query: Record<string, string | number | string[] | number[]>,
  hash?: string | null | undefined
) {
  const remainingQuery = { ...query }

  // Replace slugs, and remove them from the `query`
  let asPath = pathname.replace(/\[{1,2}(.+?)]{1,2}/g, ($0, slug: string) => {
    if (slug.startsWith('...')) slug = slug.replace('...', '')

    const value = remainingQuery[slug]!
    delete remainingQuery[slug]
    if (Array.isArray(value)) {
      return value.map(v => encodeURIComponent(v)).join('/')
    }
    return value !== undefined ? encodeURIComponent(String(value)) : ''
  })

  // Remove any trailing slashes; this can occur if there is no match for a catch-all slug ([[...slug]])
  asPath = removeTrailingSlash(asPath)

  // Ensure query values are strings
  const record = Object.entries(remainingQuery).reduce<Record<string, string>>(
    (prev, [key, value]) => {
      prev[key] = [value].join('')
      return prev
    },
    {}
  )

  // Append remaining query as a querystring, if needed:
  const qs = new URLSearchParams(record).toString()

  if (qs) asPath += `?${qs}`
  if (hash) asPath += hash

  return asPath
}

export function getBaseUrl() {
  if (typeof window !== 'undefined') return ''
  if (process.env.VERCEL_URL) return `https://${process.env.VERCEL_URL}`
  return 'http://localhost:3000'
}
