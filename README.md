# secureD

Docker Hack Day #3 Project to provide simple use of trusted images

Take the stress out of ensuring trusted images are used by providing a secureD wrapper around Docker. Set up of trust and storage of keys in a vault should be simple and ensure pull, push and build interactions with registries uses trusted capabilities

### ToDo
* Run Docker Benchmark from secureD

### Ideas:
* Wrapper around [Docker Engine](https://github.com/docker/docker) using content trust
* Implementation of Security Profiles
* Integration/Use of [dockerbench](https://github.com/docker/docker-bench-security)
* [Notary](https://github.com/docker/notary)
* [Vault](https://www.vaultproject.io/) secrets management
* [Vault](https://www.vaultproject.io/) and use of Vault [policies](https://www.vaultproject.io/docs/concepts/policies.html) for key management
* Make use of [docker-volume-vault](https://github.com/calavera/docker-volume-vault)
