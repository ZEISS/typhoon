import { ApiClient, OpenAPIConfig } from '@/generated'

const config = {}

export const newClient = () => {
  const client = new ApiClient(config)

  return client
}
