# ¿Qué son los Principios SOLID?

---

### Los Principios SOLID son un conjunto de cinco principios de diseño de software que nos guían para escribir código más limpio, modular y fácil de mantener. Fueron introducidos por Robert C. Martin, también conocido como "Uncle Bob".

---

- **_S_** - Principio de Responsabilidad Única (Single Responsibility Principle): Este principio establece que una clase debe tener una única razón para cambiar. En otras palabras, cada clase debe tener una única responsabilidad dentro del sistema.


- **_O_** - Principio de Abierto/Cerrado (Open/Closed Principle): Este principio sugiere que las entidades de software (clases, módulos, funciones, etc.) deben estar abiertas para su extensión pero cerradas para su modificación. Esto significa que deberíamos poder extender el comportamiento de una entidad sin modificar su código fuente.


- **_L_** - Principio de Sustitución de Liskov (Liskov Substitution Principle): Este principio establece que los objetos de un programa deben ser reemplazables por instancias de sus subtipos sin afectar a la ejecución del programa. En resumen, una clase derivada debe poder sustituirse por su clase base sin que el programa produzca resultados incorrectos.


- **_I_** - Principio de Segregación de Interfaces (Interface Segregation Principle): Este principio sugiere que una clase no debería estar obligada a implementar interfaces que no necesita. En lugar de tener interfaces monolíticas, es mejor tener interfaces más pequeñas y específicas.


- **_D_** - Principio de Inversión de Dependencias (Dependency Inversion Principle): Este principio establece que los módulos de alto nivel no deben depender de los módulos de bajo nivel, sino que ambos deben depender de abstracciones. Además, las abstracciones no deben depender de los detalles, sino que los detalles deben depender de las abstracciones.