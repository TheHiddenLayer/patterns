#ifndef SIMPLE_FACTORY_H
#define SIMPLE_FACTORY_H

#include <string>

class Product {
public:
    virtual ~Product() {}
    virtual std::string Operation() = 0;
};

class ProductA : public Product {
public:
    std::string Operation() override;
};

class ProductB : public Product {
public:
    std::string Operation() override;
};

Product* GetProduct(const std::string& productType);

#endif