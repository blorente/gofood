# gofood
A go application to manage food


## Features:

- Recommendations:
  - Recommend me a {mealKind} that is {predicate=high-protein, high-volume, kcal_filter} 

Technologies:
- Backend: gRPC server that extends pocketbase.
- Model: Protobuf
- CLI: Console on top of Bubbletea
- Web: WebApp on top of Svelte
- Deployment: My VPS.

## TODO:
- [ ] Filter by mealkind
- [ ] Implement the rest of the server
- [ ] Figure out filters
- [ ] Configure datadir when deploying. Create deploy file?
