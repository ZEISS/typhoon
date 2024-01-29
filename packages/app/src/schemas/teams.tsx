import { z } from 'zod'

export const TeamCreateSchema = z.object({
  name: z.string().min(3).max(128),
  slug: z.string().trim().toLowerCase().min(3).max(128).default(''),
  description: z.string().min(10).max(256).optional(),
  contactEmail: z.string().email().optional()
})

export type CreateTeamFornValues = z.infer<typeof TeamCreateSchema>
export const defaultValues: Partial<CreateTeamFornValues> = {
  name: '',
  slug: ''
}
