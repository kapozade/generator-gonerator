const chalk = require('chalk')
const directoryBuilder = require('mkdirp');
const path = require('path');
const yosay = require('yosay');

const Generator = require('yeoman-generator');

module.exports = class extends Generator {
  initializing() {
    this.destinationRoot(process.env.GOPATH || './');
  }

  prompting() {
    this.log(yosay(chalk('Welcome to yet another go microservice boilerplate!')))
    let prompts = [
      {
        type: 'input',
        name: 'appname',
        message: 'What do you want me to name your service?',
        default: 'myservice'
      },
      {
        type: 'input',
        name: 'repository',
        message: 'What is your repository?',
        default: 'github.com/username/'
      }
    ];

    return this.prompt(prompts).then(prompt => {      
      this.appname = prompt.appname.replace(/[^a-zA-Z]/g, '').toLowerCase();
      if(prompt.repository.substr(-1) != '/'){
        prompt.repository += '/';
      }
      this.repository = prompt.repository + this.appname;
    });
  }

  writing(){
    let packageDirectory = this.destinationPath('pkg');
    directoryBuilder.sync(packageDirectory);
    
    let sourceDirectory = this.destinationPath(this.repository);
    directoryBuilder.sync(sourceDirectory);

    let binaryDirectory = this.destinationPath('bin');
    directoryBuilder.sync(binaryDirectory);

    let context = {
      appname: this.appname,
      repository: this.repository
    }

    this.fs.copy(this.templatePath('gitignore.txt'), path.join(sourceDirectory, '.gitignore'));
    this.fs.copyTpl(this.templatePath('Makefile'), path.join(sourceDirectory, 'Makefile'), context);
    this.fs.copyTpl(this.templatePath('README.md'), path.join(sourceDirectory, 'README.md'), context);
    this.fs.copyTpl(this.templatePath('api/healthcheck_controller.go'), path.join(sourceDirectory, 'api/healthcheck_controller.go'), context);
    this.fs.copyTpl(this.templatePath('api/readiness_controller.go'), path.join(sourceDirectory, 'api/readiness_controller.go'), context);
    this.fs.copyTpl(this.templatePath('docs/docs.go'), path.join(sourceDirectory, 'docs/docs.go'), context);
    this.fs.copyTpl(this.templatePath('docs/swagger.json'), path.join(sourceDirectory, 'docs/swagger.json'), context);
    this.fs.copyTpl(this.templatePath('docs/swagger.yaml'), path.join(sourceDirectory, 'docs/swagger.yaml'), context);
    this.fs.copyTpl(this.templatePath('models/application_configurations.go'), path.join(sourceDirectory, 'models/application_configurations.go'), context);
    this.fs.copyTpl(this.templatePath('models/application_errors.go'), path.join(sourceDirectory, 'models/application_errors.go'), context);
    this.fs.copyTpl(this.templatePath('models/healthcheck.go'), path.join(sourceDirectory, 'models/healthcheck.go'), context);
    this.fs.copyTpl(this.templatePath('models/readiness.go'), path.join(sourceDirectory, 'models/readiness.go'), context);
    this.fs.copyTpl(this.templatePath('server/server.go'), path.join(sourceDirectory, 'server/server.go'), context);
    this.fs.copyTpl(this.templatePath('server/url_mappings.go'), path.join(sourceDirectory, 'server/url_mappings.go'), context);
    this.fs.copyTpl(this.templatePath('services/configuration_service.go'), path.join(sourceDirectory, 'services/configuration_service.go'), context);
    this.fs.copyTpl(this.templatePath('main.go'), path.join(sourceDirectory, 'main.go'), context);
    this.log(yosay('Do not forget to stay hydrated. Grab a beer now.'));
  }
}
