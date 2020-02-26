# iac-count

> iac-count by [MaibornWolff](https://www.maibornwolff.de)

## What is iac-count

`iac-count` is a tool for measuring quality of IaC. Currently it supports metrics for the following languages:

- ansible

## Usage

### Ansible

Given an ansible project (e.g. the ansible directory of [DebOps](https://github.com/debops/debops/))

You can run `iac-count` from the root of the ansible project as follows

```bash
iac-count ansible .
```

resulting in an csv output like

![iac-count Example Output](docs/images/ansible_example_csv.png "iac-count Example Output")

Using tools like [CodeCharta](https://github.com/MaibornWolff/codecharta) you can visualize the metrics. In case of [DebOps](https://github.com/debops/debops/) it looks like the following:

![CodeCharta Example](docs/images/ansible_example_codecharta.png "CodeCharta Example")

## Install / Build

Clone repository and run

```bash
go build ./...
go install ./...
```

Make sure `$GOPATH/bin` is in your `$PATH`.

## Feature request / Bug / Feedback

Have a bug, a feature request or any question? Please [open a new issue](https://github.com/MaibornWolff/iac-count/issues/new). Feedback is always welcome.

## Tool Information

- [Releases](https://github.com/MaibornWolff/iac-count/releases)
- [Contributing](CONTRIBUTING.md)
- [Code of Conduct](CODE_OF_CONDUCT.md)
- [License](LICENSE.md)