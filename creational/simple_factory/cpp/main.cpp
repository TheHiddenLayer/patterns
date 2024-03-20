#include <iostream>
#include "simple_factory.h"

int main() {
    Product* product = GetProduct("A");
    std::cout << "Product: " << product->Operation() << std::endl;
    delete product;
    return 0;
}