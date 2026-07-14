#ifndef AML_DCM_H
#define AML_DCM_H

#include "AMLMatrix33.h"
#include "AMLVector3.h"
#include <limits>

namespace AML
{
    class DCM
    {
        public:

            static const Matrix33 rotationX(double theta);
            static const Matrix33 rotationY(double theta);
            static const Matrix33 rotationZ(double theta);
    };

    bool isValidDCM(const Matrix33& dcm, double tol = std::numeric_limits<double>::epsilon());
    void normalise(Matrix33& dcm);

    // DCM Kinematic Functions
    Matrix33 integrateDCM(const Matrix33& dcm, const Matrix33& dcmRates, double dt);
    Matrix33 dcmKinematicRates_BodyRates(const Matrix33& dcm, const Vector3& bodyRates);
    Matrix33 dcmKinematicRates_WorldRates(const Matrix33& dcm, const Vector3& worldRates);

    

};

#endif //AML_DCM_H