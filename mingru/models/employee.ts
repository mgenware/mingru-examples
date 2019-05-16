import * as dd from 'dd-models';

class Employee extends dd.Table {
  id = dd.pk(dd.int()).setDBName('emp_no');
  firstName = dd.varChar(50);
  lastName = dd.varChar(50);
  gender = dd.varChar(10);
  birthDate = dd.date();
  hireDate = dd.date();
}

export default dd.table(Employee, 'employees');
